import { cn } from "@/lib/utils";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "./ui/tabs";

export interface LinkableTabProps {
	label: string;
	children: React.ReactNode;
	value?: string;
}

interface LinkableTabsProps {
	classNames?: {
		wrapper?: string;
		trigger?: string;
		content?: string;
	};
	tabs: LinkableTabProps[];
	tabSearchParam?: string;
}

const LinkableTabs: React.FC<LinkableTabsProps> = ({
	classNames,
	tabs,
	tabSearchParam,
}) => {
	const initialTab = tabSearchParam || tabs[0].value;

	return (
		<Tabs
			className={cn("p-3", classNames?.wrapper)}
			value={initialTab}
			defaultValue={tabs[0].value || tabs[0].label}
		>
			<TabsList className={"bg-transparent p-0"}>
				{tabs.map((tab, index) => (
					<TabsTrigger
						className={cn(
							initialTab === tab.value &&
								"border-b border-b-primary rounded-none",
							classNames?.trigger,
						)}
						key={`trigger-${tab.value}`}
						value={tab.value || tab.label}
						onClick={() => {
							console.log("tab.value", tab.value);
							console.log("tab.label", tab.label);
						}}
					>
						{tab.label}
					</TabsTrigger>
				))}
			</TabsList>
			{tabs.map((tab, index) => (
				<TabsContent
					className={classNames?.content}
					key={`content-${tab.value}`}
					value={tab.value || tab.label}
				>
					{tab.children}
				</TabsContent>
			))}
		</Tabs>
	);
};

export default LinkableTabs;
