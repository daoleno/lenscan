import React from "react"

interface Props {
  title: string
  metric: string
  icon: any
  color: string
}

const StatCard: React.FC<Props> = ({ title, metric, icon, color }) => {
  return (
    <div
      className={`grid h-32 grid-cols-2 rounded-2xl p-6 ${color} items-center`}
    >
      <h5 className="text-lg font-semibold text-gray-700">{title}</h5>
      <p className="text-xl font-semibold text-gray-700">{metric}</p>
    </div>
  )
}

export default StatCard
