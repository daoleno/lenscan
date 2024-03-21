import getProfileScore from "@/app/api/profiles/getProfileScore";
import { AwardIcon } from "lucide-react";

interface ProfileScoreProps {
	profileId: string;
	className?: string;
}

const ProfileScore: React.FC<ProfileScoreProps> = async ({
	profileId,
	className,
}) => {
	const profileScore = await getProfileScore(profileId);

	return (
		<div className="flex flex-col items-center justify-center p-6">
			<AwardIcon className="w-10 h-10 text-green-500" />
			<div className="text-2xl font-semibold text-foreground mt-3">
				{profileScore.score}
			</div>
			<div className="text-sm font-medium text-muted-foreground mt-2">
				Quality Score
			</div>
			<div className="text-xs font-medium text-muted-foreground mt-1">
				Generated At {new Date(profileScore.generated_at).toLocaleDateString()}
			</div>
		</div>
	);
};

export default ProfileScore;
