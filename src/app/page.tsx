"use client";
import { Search } from "@/components/search";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardFooter } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Separator } from "@/components/ui/separator";
import {
  Gem,
  MessageSquare,
  MessageSquareMore,
  ThumbsUp,
  UserRound,
} from "lucide-react";

export default function Home() {
  return (
    <div className="flex flex-col">
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
        <div className="w-1/5 p-10 flex flex-col gap-4 justify-center items-center">
          <div className="rounded-full bg-gray-800 w-[150px] h-[150px]">
            <UserRound size={150} color="white" />
          </div>
          <h3>Pasit Sri-intarasut</h3>
          <div className="flex gap-x-4">
            <h4>Follower 40</h4>
            <h4>Following 50</h4>
          </div>
        </div>
        <div className="w-3/5">
          <Card className="p-10">
            <CardContent>
              <div className="">
                <div className="flex items-center gap-4">
                  <div className="w-10 h-10 rounded-full bg-gray-800">
                    <UserRound size={40} color="white" />
                  </div>
                  <h3>Cristiano Rolando</h3>
                </div>
                <div className="mt-10">
                  Lorem ipsum dolor sit, amet consectetur adipisicing elit. Ut
                  voluptatibus perferendis tenetur, consequatur itaque rem quasi
                  fuga excepturi porro, a asperiores ducimus laboriosam at
                  magnam incidunt iste? Dicta, placeat mollitia.
                </div>
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
        </div>
      </div>
    </div>
  );
}
