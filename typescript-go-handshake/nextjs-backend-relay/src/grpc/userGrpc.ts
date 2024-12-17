import {CreateUserResponse, GetUserResponse, User} from "@/gen/user/v1/user_pb";
import userClient from "@/grpc/userClient";

async function createUser(name: string, email: string): Promise<User> {
    const response: CreateUserResponse = await userClient.createUser({
        name,
        email,
    });

    if (!response.user) {
        throw new Error("User creation failed: No user returned");
    }

    return response.user;
}

async function getUser(id: string): Promise<User> {
    const response: GetUserResponse = await userClient.getUser({
        id,
    });

    if (!response.user) {
        throw new Error(`User not found with id: ${id}`);
    }

    return response.user;
}

export {createUser, getUser}