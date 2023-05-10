import {
  Environment,
  StartDATrustingIndexingRequest,
  StreamResult,
  startDATrustingIndexing,
} from "@lens-protocol/momoka";
import prisma from "./prisma";

const stream = async (result: StreamResult) => {
  if (result.success) {
    const daResult = result.dataAvailabilityResult;
    const res = await prisma.momokaTx.upsert({
      where: { dataAvailabilityId: daResult.dataAvailabilityId },
      create: {
        proofTxId: result.proofTxId,
        success: result.success,
        dataAvailabilityId: daResult.dataAvailabilityId,
        publicationId: daResult.publicationId,
        signature: daResult.signature,
        timestampProofs: daResult.timestampProofs as any,
        chainProofs: daResult.chainProofs as any,
        timestamp: daResult.event.timestamp as any,
        type: daResult.type,
        event: daResult.event as any,
      },
      update: {},
    });
    console.log("inserted #", res.id, "proofTxId", result.proofTxId);
  } else {
    // failure reason
    console.log("reason", result.failureReason);
    // this will expose the submisson if it could be read
    console.log("submisson", result.dataAvailabilityResult);
  }
};

const request: StartDATrustingIndexingRequest = {
  environment: Environment.POLYGON,
  stream,
};

// it run forever and stream data as it comes in
startDATrustingIndexing(request);
