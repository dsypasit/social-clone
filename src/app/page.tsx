"use client";
import { Search } from "@/components/search";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Separator } from "@/components/ui/separator";
import { Gem } from "lucide-react";

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
      <div></div>
    </div>
  );
}
