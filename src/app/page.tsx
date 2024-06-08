import { Button } from "@/components/ui/button";
import Image from "next/image";

export default function Home() {
  return (
    <div className="h-screen w-screen flex flex-col justify-center items-center">
      <h2 className="text-9xl font-bold">Hello world</h2>
      <Button className="m-10">Click me</Button>
    </div>
  );
}
