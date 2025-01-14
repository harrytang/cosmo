import { useQuery } from "@tanstack/react-query";
import {
  EventMeta,
  OrganizationEventName,
} from "@wundergraph/cosmo-connect/dist/notifications/events_pb";
import { getFederatedGraphs } from "@wundergraph/cosmo-connect/dist/platform/v1/platform-PlatformService_connectquery";
import Link from "next/link";
import { useRouter } from "next/router";
import { useMemo } from "react";
import { Button } from "../ui/button";
import {
  DropdownMenu,
  DropdownMenuCheckboxItem,
  DropdownMenuContent,
  DropdownMenuTrigger,
} from "../ui/dropdown-menu";
import { Tabs, TabsList, TabsTrigger } from "../ui/tabs";
import { PartialMessage } from "@bufbuild/protobuf";
import { PiWebhooksLogo } from "react-icons/pi";
import { FaSlack } from "react-icons/fa";

export type EventsMeta = Array<PartialMessage<EventMeta>>;

type NotificationTab = "webhooks" | "integrations";

export const notificationEvents = [
  {
    id: OrganizationEventName.FEDERATED_GRAPH_SCHEMA_UPDATED,
    name: OrganizationEventName[
      OrganizationEventName.FEDERATED_GRAPH_SCHEMA_UPDATED
    ],
    label: "Federated Graph Schema Update",
    description: "An update to the schema of any federated graph",
  },
] as const;

export const SelectFederatedGraphs = ({
  meta,
  setMeta,
}: {
  meta: EventsMeta;
  setMeta: (meta: EventsMeta) => void;
}) => {
  const { data } = useQuery(getFederatedGraphs.useQuery());

  const graphIds = useMemo(() => {
    const entry = meta.find(
      (m) =>
        m.eventName === OrganizationEventName.FEDERATED_GRAPH_SCHEMA_UPDATED
    );
    if (entry?.meta?.case !== "federatedGraphSchemaUpdated") return [];
    return entry.meta.value.graphIds ?? [];
  }, [meta]);

  const onCheckedChange = (val: boolean, graphId: string) => {
    const tempMeta: EventsMeta = [...meta];
    const newGraphIds: string[] = [];

    if (val) {
      newGraphIds.push(...Array.from(new Set([...graphIds, graphId])));
    } else {
      newGraphIds.push(...graphIds.filter((g) => g !== graphId));
    }

    const entry: EventsMeta[number] = {
      eventName: OrganizationEventName.FEDERATED_GRAPH_SCHEMA_UPDATED,
      meta: {
        case: "federatedGraphSchemaUpdated",
        value: {
          graphIds: newGraphIds,
        },
      },
    };

    const idx = tempMeta.findIndex(
      (v) =>
        v.eventName === OrganizationEventName.FEDERATED_GRAPH_SCHEMA_UPDATED
    );

    if (idx === -1) {
      tempMeta.push(entry);
    } else {
      tempMeta[idx] = entry;
    }

    setMeta(tempMeta);
  };

  return (
    <DropdownMenu>
      <DropdownMenuTrigger asChild>
        <Button size="sm" variant="outline">
          {graphIds.length > 0
            ? `${graphIds.length} selected`
            : "Select graphs"}
        </Button>
      </DropdownMenuTrigger>
      <DropdownMenuContent align="start" className="">
        {data?.graphs?.map((graph) => {
          return (
            <DropdownMenuCheckboxItem
              key={graph.id}
              checked={graphIds.includes(graph.id)}
              onCheckedChange={(val) => onCheckedChange(val, graph.id)}
              onSelect={(e) => e.preventDefault()}
            >
              {graph.name}
            </DropdownMenuCheckboxItem>
          );
        })}
      </DropdownMenuContent>
    </DropdownMenu>
  );
};

export const Meta = ({
  id,
  meta,
  setMeta,
}: {
  id: OrganizationEventName;
  meta: EventsMeta;
  setMeta: (meta: EventsMeta) => void;
}) => {
  if (id == OrganizationEventName.FEDERATED_GRAPH_SCHEMA_UPDATED) {
    return <SelectFederatedGraphs meta={meta} setMeta={setMeta} />;
  }

  return null;
};

export const NotificationTabs = ({ tab }: { tab: NotificationTab }) => {
  const router = useRouter();

  return (
    <Tabs defaultValue={tab}>
      <TabsList>
        <TabsTrigger value="webhooks" asChild>
          <Link
            href={{
              pathname: `/${router.query.organizationSlug}/webhooks`,
            }}
            className="flex gap-x-[6px]"
          >
            <PiWebhooksLogo />
            Webhooks
          </Link>
        </TabsTrigger>
        <TabsTrigger value="integrations" asChild>
          <Link
            href={{
              pathname: `/${router.query.organizationSlug}/integrations`,
            }}
            className="flex gap-x-2"
          >
            <FaSlack />
            Slack Integration
          </Link>
        </TabsTrigger>
      </TabsList>
    </Tabs>
  );
};
