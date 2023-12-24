import Image from "next/image"

import { getLenny } from "@/lib/lenny"
import { cn } from "@/lib/utils"

import { Card } from "./ui/card"

interface Attribute {
  display_type?: string
  trait_type: string
  value: string
}

interface Lenny {
  image: string
  attributes: Attribute[]
}

interface LennyCardProps {
  profileId: string
  className?: string
}

const LennyCard: React.FC<LennyCardProps> = async ({
  profileId,
  className,
}) => {
  const lenny: Lenny = await getLenny(profileId)

  return (
    <Card className={cn(className, "p-3")}>
      <div className="flex flex-col items-center justify-between gap-7 sm:flex-row">
        <Image
          src={lenny.image}
          alt="Lenny Avatar"
          width={300}
          height={300}
          className="h-32 w-32 rounded-xl object-cover sm:h-52 sm:w-52"
        />
        <div className="grid w-auto grid-cols-2 gap-2.5 text-xs">
          {lenny.attributes
            .sort((a, b) => a.trait_type.localeCompare(b.trait_type))
            .map((attr, index) => (
              <div key={index} className="flex gap-1">
                <div className="text-muted-foreground">{attr.trait_type}:</div>
                <div className="font-semibold capitalize text-foreground">
                  {attr.value}
                </div>
              </div>
            ))}
        </div>
      </div>
    </Card>
  )
}

export default LennyCard
