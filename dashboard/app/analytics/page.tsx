import { AreaChart, BarChart, DonutChart, LineChart } from "@tremor/react"

import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"

import { getPopularHashtags } from "../api/analystics/getPopularHashtags"
import { getPublicationReactionsOvertime } from "../api/analystics/getPublicationReactionsOverTime"
import { getPublicationTypesDistribution } from "../api/analystics/getPublicationTypesDistribution"
import { getUesrActivities } from "../api/analystics/getUserActivities"

export default async function Page() {
  const userActivity = await getUesrActivities()
  const hashtags = await getPopularHashtags()
  const publicationTypesDeistribution = await getPublicationTypesDistribution()
  //   const userInteractionNetwork = await duckdb.all(
  //     `SELECT
  //     pr1.profile_id AS source_user,
  //     pr2.profile_id AS target_user,
  //     COUNT(*) AS interaction_count
  // FROM
  //     publication_mention pm
  // JOIN
  //     publication_record pr1 ON pm.publication_id = pr1.id
  // JOIN
  //     publication_record pr2 ON pm.mentioned_profile_id = pr2.profile_id
  // GROUP BY
  //     pr1.profile_id, pr2.profile_id;
  //     `
  //   )
  const publicationReactionsOverTime = await getPublicationReactionsOvertime()
  console.log(publicationReactionsOverTime)

  return (
    <div className="grid grid-cols-1 gap-4 md:grid-cols-2">
      <Card>
        <CardHeader>
          <CardTitle>Users Activity</CardTitle>
        </CardHeader>
        <CardContent>
          <LineChart
            data={userActivity}
            categories={["count"]}
            index="day"
            showAnimation
          />
        </CardContent>
      </Card>
      <Card>
        <CardHeader>
          <CardTitle>Popular Hashtags</CardTitle>
        </CardHeader>
        <CardContent>
          <BarChart
            data={hashtags}
            categories={["count"]}
            index="hashtag"
            showAnimation
          />
        </CardContent>
      </Card>
      <Card>
        <CardHeader>
          <CardTitle>Publication Types Distribution</CardTitle>
        </CardHeader>
        <CardContent>
          <DonutChart
            data={publicationTypesDeistribution}
            category="count"
            index="publication_type"
            showAnimation
          />
        </CardContent>
      </Card>
      <Card>
        <CardHeader>
          <CardTitle>Publication Reactions Over Time</CardTitle>
        </CardHeader>
        <CardContent>
          <AreaChart
            className="mt-4 h-72"
            data={publicationReactionsOverTime}
            index="date"
            categories={[
              "UPVOTE",
              "DOWNVOTE",
              // ... other reaction types
            ]}
            stack
          />
        </CardContent>
      </Card>
    </div>
  )
}
