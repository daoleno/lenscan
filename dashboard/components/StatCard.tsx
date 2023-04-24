import { ChatBubbleLeftEllipsisIcon } from "@heroicons/react/24/outline";
import React from "react";

interface Props {
  title: string;
  metric: string;
  icon: any;
  color: string;
}

const StatCard: React.FC<Props> = ({ title, metric, icon, color }) => {
  return (
    <div
      className={`grid grid-cols-2 p-6 rounded-2xl h-32 ${color} items-center`}
    >
      <h5 className="text-lg font-semibold text-gray-700">{title}</h5>
      <ChatBubbleLeftEllipsisIcon className="w-8 h-8" />
      <p className="text-xl font-semibold text-gray-700">{metric}</p>
    </div>
  );
};

export default StatCard;
