import {UserService} from '@/gen/user/v1/user_pb';
import {createClient} from "@connectrpc/connect";
import {createGrpcTransport} from "@connectrpc/connect-node";

const transport = createGrpcTransport({
    baseUrl: "http://localhost:50551",
})

export default createClient(UserService, transport);
