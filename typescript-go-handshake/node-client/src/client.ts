import {
    User,
    CreateUserRequest,
    CreateUserResponse,
    GetUserRequest,
    GetUserResponse,
    UserService,
} from "./gen/user/v1/user_pb";
import {createClient} from "@connectrpc/connect";
import {createGrpcTransport} from "@connectrpc/connect-node";


function creategRPCClient() {
    const transport = createGrpcTransport({
        baseUrl: "http://localhost:50551",
    });
    return createClient(UserService, transport);
}
const client = creategRPCClient();

async function createUser(name: string, email: string): Promise<User> {
    const response: CreateUserResponse = await client.createUser({
        name,
        email,
    });

    if (!response.user) {
        throw new Error("User creation failed: No user returned");
    }

    return response.user;
}

async function getUser(id: string): Promise<User> {
    const response: GetUserResponse = await client.getUser({
        id,
    });

    if (!response.user) {
        throw new Error(`User not found with id: ${id}`);
    }

    return response.user;
}

async function main() {
    try {
        console.log("Creating new user...");
        const newUser = await createUser("John Doe", "john@example.com");
        console.log("Created user:", {
            id: newUser.id,
            name: newUser.name,
            email: newUser.email,
        });

        console.log("\nFetching user details...");
        const fetchedUser = await getUser(newUser.id);
        console.log("Retrieved user:", {
            id: fetchedUser.id,
            name: fetchedUser.name,
            email: fetchedUser.email,
        });
    } catch (error) {
        if (error instanceof Error) {
            console.error("Application error:", error.message);
        } else {
            console.error("Unknown error occurred:", error);
        }
        process.exit(1);
    }
}

process.on("SIGINT", () => {
    console.log("\nGracefully shutting down...");
    process.exit(0);
});

await main()