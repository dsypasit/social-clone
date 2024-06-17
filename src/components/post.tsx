import { UserRound, ThumbsUp, MessageSquare } from "lucide-react";
import { Card, CardContent, CardFooter } from "./ui/card";
import { IPost } from "@/types/types";

interface PostProps {
  post: IPost; // Prop that expects a Post object
}

export const Post: React.FC<PostProps> = ({ post }) => {
  return (
    <Card className="p-10">
      <CardContent>
        <div className="">
          <div className="flex items-center gap-4">
            <div className="w-10 h-10 rounded-full bg-gray-800">
              <UserRound size={40} color="white" />
            </div>
            <h3>{post.username}</h3>
          </div>
          <div className="mt-10">{post.content}</div>
        </div>
      </CardContent>
      <CardFooter>
        <div className="w-full">
          <div className="flex gap-x-4">
            <div className="flex cursor-pointer gap-x-2">
              <ThumbsUp />
              20
            </div>

            {/* TODO: use dialog to toggle comment */}
            <div className="flex cursor-pointer gap-x-2">
              <MessageSquare />
              20
            </div>
          </div>
        </div>
      </CardFooter>
    </Card>
  );
};
