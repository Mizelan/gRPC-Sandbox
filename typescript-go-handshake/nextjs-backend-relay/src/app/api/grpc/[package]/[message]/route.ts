import {NextRequest,NextResponse} from "next/server";
import {fromBinary, toBinary} from "@bufbuild/protobuf";
import {encodeEnvelope} from "@connectrpc/connect/protocol";
import userClient from "@/grpc/userClient";
import userBackend from "@/grpc/userBackend";

const userProtobufPath = '@/gen/user/v1/user_pb';

export async function POST(request: NextRequest,{ params }: { params: Promise<{ package: string, message: string }>}) {
    const pkg = (await params).package;
    const message = (await params).message;
    const methodName = message.charAt(0).toLowerCase() + message.slice(1);
    const requestSchema = await import(userProtobufPath).then(module => module[`${message}RequestSchema`]);
    const responseSchema = await import(userProtobufPath).then(module => module[`${message}ResponseSchema`]);
    console.log(`Received request for ${pkg}.${message} with method ${methodName}`);

    // decode request data
    const requestBinary = await request.arrayBuffer();
    const requestBytes = new Uint8Array(requestBinary);
    const compressedFlag = requestBytes[0] & 0x80;
    const messageType = requestBytes[0] & 0x7f;
    const decodedLength = requestBytes[1] << 24 | requestBytes[2] << 16 | requestBytes[3] << 8 | requestBytes[4];
    const messageBytes = requestBytes.slice(5, 5 + decodedLength);
    const requestData = fromBinary(requestSchema, messageBytes);
    console.log("Decoded request data:", requestData);
    console.log("Compressed flag:", compressedFlag);
    console.log("Message type:", messageType);

    // call service method dynamically
    const response = await userBackend[methodName](requestData);
    const binaryResponse = toBinary(responseSchema, response);
    const messageFrame = encodeEnvelope(0, binaryResponse);
    const trailerFrame = encodeEnvelope(128, new Uint8Array(0));
    const fullResponse = new Uint8Array(messageFrame.length + trailerFrame.length);
    fullResponse.set(messageFrame, 0);
    fullResponse.set(trailerFrame, messageFrame.length);

    return new Response(fullResponse, {
        headers: {
            'content-type': 'application/grpc-web+proto',
            'grpc-status': '0',
            'grpc-encoding': 'identity'
        }
    });
}