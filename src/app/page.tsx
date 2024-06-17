"use client";
import { CreatedPost } from "@/components/createdPost";
import { Post } from "@/components/post";
import { Search } from "@/components/search";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardFooter } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Separator } from "@/components/ui/separator";
import { Textarea } from "@/components/ui/textarea";
import apiClient from "@/lib/apiClient";
import { IPost } from "@/types/types";
import {
  Gem,
  MessageSquare,
  MessageSquareMore,
  ThumbsUp,
  UserRound,
} from "lucide-react";
import { useEffect, useState } from "react";

export default function Home() {
  const [posts, setPosts] = useState<IPost[]>([]); // State to store fetched posts

  useEffect(() => {
    const fetchPosts = async () => {
      try {
        const response = await apiClient.get("/post"); // Send GET request to /posts
        setPosts(response.data); // Update state with fetched posts
      } catch (error) {
        console.error("Error fetching posts:", error); // Handle errors
        // You can display an error message to the user here
      }
    };

    fetchPosts(); // Call the function to fetch posts on component mount
  }, []); // Empty dependency array ensures fetching only once on mount

  // Function to update the posts state (passed as a prop to CreatedPost)
  const updatePosts = (newPost: IPost[]) => {
    setPosts(newPost); // Update state with the new post
  };

  return (
    <div className="flex flex-col">
      {/* NOTE: Header */}
      <div className="relative w-full p-5 flex items-center justify-center">
        <div className="absolute top-5 left-5 flex gap-4">
          <Gem className="" />
          <h1>social-clone</h1>
        </div>
        <div className="w-2/5 z-20">
          <Search />
        </div>
      </div>
      <div className="flex px-10">
        {/* NOTE: Left section */}
        <div className="w-1/5 p-10 flex flex-col gap-4 items-center">
          <div className="rounded-full bg-gray-800 w-[150px] h-[150px]">
            <UserRound size={150} color="white" />
          </div>
          <h3>Pasit Sri-intarasut</h3>
          <div className="flex gap-x-4">
            <h4>Follower 40</h4>
            <h4>Following 50</h4>
          </div>
        </div>
        {/* NOTE: Right section */}
        <div className="w-3/5">
          {/* NOTE: Create Post section */}
          <CreatedPost updatePosts={updatePosts} />
          {/* NOTE: Posts list */}
          {posts.map((post, i) => (
            <Post key={i} post={post} /> // Pass each post object to Post component
          ))}
          {posts.length === 0 ? <p>No posts found.</p> : ""}
        </div>
      </div>
    </div>
  );
}
