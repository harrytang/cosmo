import { federateSubgraphs, RootTypeFieldData, Subgraph, unresolvableFieldError } from '../src';
import { parse } from 'graphql';
import { describe, expect, test } from 'vitest';
import {
  documentNodeToNormalizedString,
  normalizeString,
  versionOnePersistedBaseSchema,
  versionTwoPersistedBaseSchema,
} from './utils/utils';

describe('Query federation tests', () => {
  test('that shared queries that return a nested type that is only resolvable over multiple subgraphs are valid', () => {
    const { errors, federationResult } = federateSubgraphs([subgraphA, subgraphB]);
    expect(errors).toBeUndefined();
    const federatedGraph = federationResult!.federatedGraphAST;
    expect(documentNodeToNormalizedString(federatedGraph)).toBe(
      normalizeString(
        versionTwoPersistedBaseSchema +
          `
      type Query {
        query: Nested
      }

      type Nested {
        nest: Nested2
      }

      type Nested2 {
        nest: Nested3
      }

      type Nested3 {
        nest: Nested4
      }

      type Nested4 {
        name: String
        age: Int
      }
    `,
      ),
    );
  });

  test('that unshared queries that return a nested type that cannot be resolved in a single subgraph returns an error', () => {
    const rootTypeFieldData: RootTypeFieldData = {
      fieldName: 'query',
      fieldTypeNodeString: 'Nested',
      path: 'Query.query',
      subgraphs: new Set<string>(['subgraph-b']),
      typeName: 'Query',
    };
    const { errors } = federateSubgraphs([subgraphB, subgraphC]);
    expect(errors).toBeDefined();
    expect(errors).toHaveLength(1);
    expect(errors![0]).toStrictEqual(
      unresolvableFieldError(
        rootTypeFieldData,
        'name',
        ['subgraph-c'],
        'Query.query.nest.nest.nest.name',
        'Nested4',
      ),
    );
  });

  test('that unresolvable fields return an error', () => {
    const rootTypeFieldData: RootTypeFieldData = {
      fieldName: 'friend',
      fieldTypeNodeString: 'Friend',
      path: 'Query.friend',
      subgraphs: new Set<string>(['subgraph-d']),
      typeName: 'Query',
    };
    const { errors } = federateSubgraphs([subgraphD, subgraphF]);
    expect(errors).toBeDefined();
    expect(errors).toHaveLength(1);
    expect(errors![0]).toStrictEqual(
      unresolvableFieldError(
        rootTypeFieldData,
        'age',
        ['subgraph-f'],
        'Query.friend.age',
        'Friend',
      ),
    );
  });

  test('that unresolvable fields that are the first fields to be added still return an error', () => {
    const rootTypeFieldData: RootTypeFieldData = {
      fieldName: 'friend',
      fieldTypeNodeString: 'Friend',
      path: 'Query.friend',
      subgraphs: new Set<string>(['subgraph-d']),
      typeName: 'Query',
    };
    const { errors } = federateSubgraphs([subgraphF, subgraphD]);
    expect(errors).toBeDefined();
    expect(errors).toHaveLength(1);
    expect(errors![0]).toStrictEqual(
      unresolvableFieldError(
        rootTypeFieldData,
        'age',
        ['subgraph-f'],
        'Query.friend.age',
        'Friend',
      ),
    );
  });

  test('that multiple unresolved fields return an error for each', () => {
    const rootTypeFieldData: RootTypeFieldData = {
      fieldName: 'friend',
      fieldTypeNodeString: 'Friend',
      path: 'Query.friend',
      subgraphs: new Set<string>(['subgraph-d']),
      typeName: 'Query',
    };
    const { errors } = federateSubgraphs([subgraphD, subgraphF, subgraphG]);
    expect(errors).toBeDefined();
    expect(errors).toHaveLength(2);
    expect(errors![0]).toStrictEqual(
      unresolvableFieldError(
        rootTypeFieldData,
        'age',
        ['subgraph-f'],
        'Query.friend.age',
        'Friend',
      ),
    );
    expect(errors![1]).toStrictEqual(
      unresolvableFieldError(
        rootTypeFieldData,
        'hobbies',
        ['subgraph-g'],
        'Query.friend.hobbies',
        'Friend',
      ),
    );
  });

  test('that shared queries that return a type that is only resolvable over multiple subgraphs are valid', () => {
    const { errors, federationResult } = federateSubgraphs([subgraphD, subgraphE]);
    expect(errors).toBeUndefined();
    const federatedGraph = federationResult!.federatedGraphAST;
    expect(documentNodeToNormalizedString(federatedGraph)).toBe(
      normalizeString(
        versionTwoPersistedBaseSchema +
          `
      type Query {
        friend: Friend
      }

      type Friend {
        name: String!
        age: Int!
      }
    `,
      ),
    );
  });

  test('that shared queries that return an interface that is only resolvable over multiple subgraphs are valid', () => {
    const { errors, federationResult } = federateSubgraphs([subgraphH, subgraphI]);
    expect(errors).toBeUndefined();
    const federatedGraph = federationResult!.federatedGraphAST;
    expect(documentNodeToNormalizedString(federatedGraph)).toBe(
      normalizeString(
        versionOnePersistedBaseSchema +
          `
      interface Human {
        name: String!
        age: Int!
      }
      
      type Query {
        humans: [Human]
      }

      type Friend implements Human {
        name: String!
        age: Int!
      }
    `,
      ),
    );
  });

  test('that queries that return interfaces whose constituent types are unresolvable return an error', () => {
    const rootTypeFieldData: RootTypeFieldData = {
      fieldName: 'humans',
      fieldTypeNodeString: '[Human]',
      path: 'Query.humans',
      subgraphs: new Set<string>(['subgraph-i']),
      typeName: 'Query',
    };
    const result = federateSubgraphs([subgraphI, subgraphJ]);
    expect(result.errors).toBeDefined();
    expect(result.errors).toHaveLength(1);
    expect(result.errors![0]).toStrictEqual(
      unresolvableFieldError(
        rootTypeFieldData,
        'name',
        ['subgraph-j'],
        'Query.humans ... on Friend name',
        'Friend',
      ),
    );
  });

  test('that queries that return nested interfaces whose constituent types are unresolvable return an error', () => {
    const rootTypeFieldData: RootTypeFieldData = {
      fieldName: 'humans',
      fieldTypeNodeString: '[Human]',
      path: 'Query.humans',
      subgraphs: new Set<string>(['subgraph-k']),
      typeName: 'Query',
    };
    const { errors } = federateSubgraphs([subgraphK, subgraphL]);
    expect(errors).toBeDefined();
    expect(errors).toHaveLength(1);
    expect(errors![0]).toStrictEqual(
      unresolvableFieldError(
        rootTypeFieldData,
        'age',
        ['subgraph-l'],
        'Query.humans ... on Friend pets ... on Cat age',
        'Cat',
      ),
    );
  });

  test('that shared queries that return a union that is only resolvable over multiple subgraphs are valid', () => {
    const { errors, federationResult } = federateSubgraphs([subgraphM, subgraphN]);
    expect(errors).toBeUndefined();
    const federatedGraph = federationResult!.federatedGraphAST;
    expect(documentNodeToNormalizedString(federatedGraph)).toBe(
      normalizeString(
        versionOnePersistedBaseSchema +
          `
      union Human = Friend | Enemy
      
      type Query {
        humans: [Human]
      }

      type Friend {
        name: String!
      }
      
      type Enemy {
        name: String!
      }
    `,
      ),
    );
  });

  test('that queries that return unions whose constituent types are unresolvable return an error', () => {
    const rootTypeFieldData: RootTypeFieldData = {
      fieldName: 'humans',
      fieldTypeNodeString: '[Human]',
      path: 'Query.humans',
      subgraphs: new Set<string>(['subgraph-o']),
      typeName: 'Query',
    };
    const result = federateSubgraphs([subgraphO, subgraphP]);
    expect(result.errors).toBeDefined();
    expect(result.errors).toHaveLength(1);
    expect(result.errors![0]).toStrictEqual(
      unresolvableFieldError(
        rootTypeFieldData,
        'age',
        ['subgraph-p'],
        'Query.humans ... on Enemy age',
        'Enemy',
      ),
    );
  });

  test('that an entity ancestor provides access to an otherwise unreachable field', () => {
    const { errors, federationResult } = federateSubgraphs([subgraphQ, subgraphR]);
    expect(errors).toBeUndefined();
    expect(documentNodeToNormalizedString(federationResult!.federatedGraphAST)).toBe(normalizeString(
      versionOnePersistedBaseSchema + `
        type Query {
          entity: SometimesEntity!
        }
        
        type SometimesEntity {
            id: ID!
            object: Object!
        }
        
        type Object {
            nestedObject: NestedObject!
        }
        
        type NestedObject {
            name: String!
            age: Int!
        }
    `));
  });

  test('that a nested self-referential type does not create an infinite validation loop', () => {
    const { errors, federationResult } = federateSubgraphs([subgraphS, subgraphD]);
    expect(errors).toBeUndefined();
    expect(documentNodeToNormalizedString(federationResult!.federatedGraphAST)).toBe(normalizeString(
      versionTwoPersistedBaseSchema + `
        type Query {
          object: Object!
          friend: Friend
        }
        
        type Object {
          nestedObject: NestedObject!
        }
        
        type NestedObject {
          object: Object!
        }
        
        type Friend {
          name: String!
        }
    `));
  });

  test('that unreachable interface implementations do not return an error', () => {
    const { errors, federationResult } = federateSubgraphs([subgraphT, subgraphU]);
    expect(errors).toBeUndefined();
    expect(documentNodeToNormalizedString(federationResult!.federatedGraphAST)).toBe(normalizeString(
      versionOnePersistedBaseSchema + `
        interface Interface {
          field: String!
        }

        type Query {
          query: Interface!
        }

        type Object implements Interface {
          field: String!
        }

        type OtherObject implements Interface {
          field: String!
        }
    `));
  });
});

const subgraphA: Subgraph = {
  name: 'subgraph-a',
  url: '',
  definitions: parse(`
    type Query {
      query: Nested @shareable
    }

    type Nested @shareable {
      nest: Nested2
    }

    type Nested2 @shareable {
      nest: Nested3
    }

    type Nested3 @shareable {
      nest: Nested4
    }

    type Nested4 {
      name: String
    }
  `),
};

const subgraphB: Subgraph = {
  name: 'subgraph-b',
  url: '',
  definitions: parse(`
    type Query {
      query: Nested @shareable
    }

    type Nested @shareable {
      nest: Nested2
    }

    type Nested2 @shareable {
      nest: Nested3
    }

    type Nested3 @shareable {
      nest: Nested4
    }

    type Nested4 {
      age: Int
    }
  `),
};

const subgraphC: Subgraph = {
  name: 'subgraph-c',
  url: '',
  definitions: parse(`
    type Nested @shareable {
      nest: Nested2
    }

    type Nested2 @shareable {
      nest: Nested3
    }

    type Nested3 @shareable {
      nest: Nested4
    }

    type Nested4 {
      name: String
    }
  `),
};

const subgraphD: Subgraph = {
  name: 'subgraph-d',
  url: '',
  definitions: parse(`
    type Query {
      friend: Friend @shareable
    }

    type Friend {
      name: String!
    }
  `),
};

const subgraphE: Subgraph = {
  name: 'subgraph-e',
  url: '',
  definitions: parse(`
    type Query {
      friend: Friend @shareable
    }

    type Friend {
      age: Int!
    }
  `),
};

const subgraphF: Subgraph = {
  name: 'subgraph-f',
  url: '',
  definitions: parse(`
    type Friend {
      age: Int!
    }
  `),
};

const subgraphG: Subgraph = {
  name: 'subgraph-g',
  url: '',
  definitions: parse(`
    type Friend {
      hobbies: [String!]!
    }
  `),
};

const subgraphH: Subgraph = {
  name: 'subgraph-h',
  url: '',
  definitions: parse(`
    type Query {
      humans: [Human]
    }
    
    interface Human {
      name: String!
    }
    
    type Friend implements Human {
      name: String!
    }
  `),
};

const subgraphI: Subgraph = {
  name: 'subgraph-i',
  url: '',
  definitions: parse(`
    type Query {
      humans: [Human]
    }
    
    interface Human {
      age: Int!
    }
    
    type Friend implements Human {
      age: Int!
    }
  `),
};

const subgraphJ: Subgraph = {
  name: 'subgraph-j',
  url: '',
  definitions: parse(`
    interface Human {
      name: String!
    }
    
    type Friend implements Human {
      name: String!
    }
  `),
};

const subgraphK: Subgraph = {
  name: 'subgraph-k',
  url: '',
  definitions: parse(`
    type Query {
      humans: [Human]
    }
    
    interface Human {
      name: String!
      pets: [Pet]
    }
    
    interface Pet {
      name: String!
    }
    
    type Cat implements Pet {
      name: String!
    }
    
    type Friend implements Human {
      name: String!
      pets: [Pet]
    }
  `),
};

const subgraphL: Subgraph = {
  name: 'subgraph-l',
  url: '',
  definitions: parse(`
    interface Human {
      name: String!
      pets: [Pet]
    }
    
    interface Pet {
      age: Int!
    }
    
    type Cat implements Pet {
      age: Int!
    }
    
    type Friend implements Human {
      name: String!
      pets: [Pet]
    }
  `),
};

const subgraphM: Subgraph = {
  name: 'subgraph-m',
  url: '',
  definitions: parse(`
    type Query {
      humans: [Human]
    }
    
    union Human = Friend
    
    type Friend {
      name: String!
    }
  `),
};

const subgraphN: Subgraph = {
  name: 'subgraph-n',
  url: '',
  definitions: parse(`
    type Query {
      humans: [Human]
    }
    
    union Human = Enemy
    
    type Enemy {
      name: String!
    }
  `),
};

const subgraphO: Subgraph = {
  name: 'subgraph-o',
  url: '',
  definitions: parse(`
    type Query {
      humans: [Human]
    }
    
    union Human = Friend | Enemy
    
    type Friend {
      name: String!
    }
    
    type Enemy {
      name: String!
    }
  `),
};

const subgraphP: Subgraph = {
  name: 'subgraph-p',
  url: '',
  definitions: parse(`
    union Human = Enemy
    
    type Enemy {
      age: Int!
    }
  `),
};

const subgraphQ = {
  name: 'subgraph-q',
  url: '',
  definitions: parse(`
    type Query {
      entity: SometimesEntity!
    }
    
    type SometimesEntity {
        id: ID!
        object: Object!
    }
    
    type Object {
        nestedObject: NestedObject!
    }
    
    type NestedObject {
        name: String!
    }
  `),
};

const subgraphR = {
  name: 'subgraph-r',
  url: '',
  definitions: parse(`
    type SometimesEntity @key(fields: "id") {
        id: ID!
        object: Object!
    }
    
    type Object {
        nestedObject: NestedObject!
    }
    
    type NestedObject {
        age: Int!
    }
  `),
};

const subgraphS = {
  name: 'subgraph-s',
  url: '',
  definitions: parse(`
    type Query {
        object: Object!
    }
    
    type Object {
        nestedObject: NestedObject!
    }
    
    type NestedObject {
        object: Object!
    }
  `),
};

const subgraphT = {
  name: 'subgraph-t',
  url: '',
  definitions: parse(`
    type Query {
     query: Interface!
    }
    
    interface Interface {
     field: String!
    }
    
    type Object implements Interface {
     field: String!
    }
  `),
};

const subgraphU = {
  name: 'subgraph-u',
  url: '',
  definitions: parse(`
    interface Interface {
     field: String!
    }
    
    type OtherObject implements Interface {
     field: String!
    }
  `),
};

