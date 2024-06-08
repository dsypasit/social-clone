"use client";
import { TypeAnimation } from "react-type-animation";

export default function authLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <div className="flex">
      <div className="relative w-1/2 bg-cover bg-commu bg-center flex items-center">
        <div className="p-10 z-10 tracking-widest">
          <TypeAnimation
            className="text-8xl text-white "
            sequence={[
              "Welcome to Social Clone",
              5000,
              "Welcome to reconnect and discover",
              5000,
            ]}
            speed={50}
            repeat={Infinity}
          />
        </div>
        <div className="absolute insert-0 w-full h-full bg-black opacity-80"></div>
      </div>
      {children}
    </div>
  );
}
