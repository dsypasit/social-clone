"use client";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";

import { useForm } from "react-hook-form";

import { yupResolver } from "@hookform/resolvers/yup";
import * as yup from "yup";
import { loginSchema } from "@/app/(auth)/login/validate";
import Link from "next/link";

type FormData = yup.InferType<typeof loginSchema>;

export default function Login() {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<FormData>({
    resolver: yupResolver(loginSchema),
  });
  const onSubmit = (data: FormData) => console.log(data);

  return (
    <div className="h-screen w-1/2 px-20 flex flex-col justify-center items-center">
      <h2 className="text-4xl font-bold mb-20">Login</h2>
      <form className="flex flex-col w-3/5" onSubmit={handleSubmit(onSubmit)}>
        <label>Username</label>
        <Input
          {...register("username")}
          type="text"
          placeholder="username"
          className="mt-4"
        />
        <p className="text-destructive text-sm">{errors.username?.message}</p>
        <label className="mt-4">Password</label>
        <Input
          {...register("password")}
          type="password"
          placeholder="password"
          className="mt-4 tracking-widest"
        />
        <p className="text-destructive text-sm">{errors.password?.message}</p>
        <p className="mt-4 text-sm text-gray-600">
          <Link
            className=" underline hover:text-red-300 duration-200"
            href="/signup"
          >
            Sign up
          </Link>{" "}
          if you don't have an account yet.
        </p>
        <Button className="mt-4">Click me</Button>
      </form>
    </div>
  );
}
