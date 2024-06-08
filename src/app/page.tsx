import { Button } from "@/components/ui/button";

export default function Home() {
  return (
    <div className="h-screen w-screen flex flex-col justify-center items-center">
      <h2 className="text-9xl font-bold">Hello world</h2>
      <Button>Click me</Button>
    </div>
  );
}
