import {
  Tooltip,
  TooltipContent,
  TooltipProvider,
  TooltipTrigger,
} from "@/components/ui/tooltip";

export default function SimpleTooltip({
  tip,
  children,
}: {
  tip: string | undefined;
  children: React.ReactNode;
}) {
  return (
    <TooltipProvider>
      <Tooltip>
        <TooltipTrigger>{children}</TooltipTrigger>
        <TooltipContent>{tip ?? "No tip provided."}</TooltipContent>
      </Tooltip>
    </TooltipProvider>
  );
}
