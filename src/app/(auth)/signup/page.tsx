"use client";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";

import { useForm } from "react-hook-form";

import { yupResolver } from "@hookform/resolvers/yup";
import * as yup from "yup";
import { loginSchema } from "@/app/(auth)/signup/validate";
import Link from "next/link";

type FormData = yup.InferType<typeof loginSchema>;

export default function Signup() {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<FormData>({
    resolver: yupResolver(loginSchema),
    mode: "all",
  });
  const onSubmit = (data: FormData) => console.log(data);

  return (
    <div className="h-screen w-1/2 px-20 flex flex-col justify-center items-center">
      <h2 className="text-4xl font-bold mb-20">Sign Up</h2>
      <form className="flex flex-col w-3/5" onSubmit={handleSubmit(onSubmit)}>
        <label>Username</label>
        <Input
          {...register("username")}
          type="text"
          placeholder="username"
          className="mt-4"
        />
        <p className="mt-1 text-destructive text-sm">
          {errors.username?.message}
        </p>
        <label className="mt-4">Email</label>
        <Input
          {...register("email")}
          type="text"
          placeholder="email"
          className="mt-4"
        />
        <p className="mt-1 text-destructive text-sm">{errors.email?.message}</p>
        <label className="mt-4">Password</label>
        <Input
          {...register("password")}
          type="password"
          placeholder="password"
          className="mt-4 tracking-widest"
        />
        <p className="mt-1 text-destructive text-sm">
          {errors.password?.message}
        </p>
        <Button className="mt-4">Click me</Button>
      </form>
    </div>
  );
}
