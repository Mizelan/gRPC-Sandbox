import {UserService} from '@/gen/user/v1/user_pb';
import {createClient} from "@connectrpc/connect";
import {createGrpcWebTransport} from "@connectrpc/connect-web";

const transport = createGrpcWebTransport({
    baseUrl: "/api/grpc",
    fetch: async (input: RequestInfo | URL, init?: RequestInit) => {
        return fetch(input, {
            ...init,
            method: "POST",
            credentials: "include",
            mode: "cors",
            body: init?.body
        });
    },
});

export default createClient(UserService, transport);
