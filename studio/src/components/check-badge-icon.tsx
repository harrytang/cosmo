import { CheckCircledIcon, CrossCircledIcon } from "@radix-ui/react-icons";
import { Badge } from "./ui/badge";

const getCheckBadge = (
  isBreaking: boolean,
  isComposable: boolean,
  isForced: boolean
) => {
  if (isForced) {
    return <Badge variant="outline">FORCED</Badge>;
  }

  return isComposable && !isBreaking ? (
    <Badge variant="success">PASSED</Badge>
  ) : (
    <Badge variant="destructive">FAILED</Badge>
  );
};

const getCheckIcon = (check: boolean) => {
  if (check) {
    return (
      <div className="flex justify-center">
        <CheckCircledIcon className="h-4 w-4 text-success" />
      </div>
    );
  }
  return (
    <div className="flex justify-center">
      <CrossCircledIcon className="h-4 w-4 text-destructive" />
    </div>
  );
};

export { getCheckBadge, getCheckIcon };
