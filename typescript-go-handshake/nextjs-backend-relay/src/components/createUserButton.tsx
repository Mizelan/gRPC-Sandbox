'use client';

import {useCallback, useState} from "react";
import {createUser} from "@/grpc/userGrpc";
import {GetUserResponse, User} from "@/gen/user/v1/user_pb";

export default function () {
    const [userResponse, setUserResponse] = useState<User|string>(null);

    const handleClick = useCallback(async () => {
        try {
            const response = await createUser("Test User", "user@test.com");
            console.log("Created user:", response);
            setUserResponse(response);
        } catch (error) {
            // Handle errors by showing them to the user
            console.error("Error creating user:", error);
            setUserResponse(error.message);
        }
    }, []);

    const getDisplayText = () => {
        if (!userResponse) return null;
        if (typeof userResponse === 'string') {
            return `Error: ${userResponse}`;
        }

        return `Created user: ${userResponse.name} (ID: ${userResponse.id})`;
    };


    return (
        <div className="flex flex-col items-center gap-4">
            <a
                className="rounded-full border border-solid border-transparent transition-colors flex items-center justify-center bg-foreground text-background gap-2 hover:bg-[#383838] dark:hover:bg-[#ccc] text-sm sm:text-base h-10 sm:h-12 px-4 sm:px-5"
                onClick={handleClick}
            >
                Create User
            </a>

            {userResponse && (
                <p className="text-center text-sm sm:text-base">
                    {getDisplayText()}
                </p>
            )}
        </div>
    )
}