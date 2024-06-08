"use client";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";

import { useForm } from "react-hook-form";

import { yupResolver } from "@hookform/resolvers/yup";
import * as yup from "yup";
import { loginSchema } from "@/app/(auth)/signup/validate";
import { useRouter } from "next/navigation";
import Link from "next/link";
import { signupService } from "@/lib/services/authService";
import { Loader2 } from "lucide-react";
import { useState } from "react";
import { inspect } from "util";

type FormData = yup.InferType<typeof loginSchema>;

export default function Signup() {
  const [isSubmitting, setIsSubmitting] = useState(false);
  const [error, setError] = useState("");
  const router = useRouter();
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<FormData>({
    resolver: yupResolver(loginSchema),
    mode: "all",
  });
  const onSubmit = async (data: FormData) => {
    setIsSubmitting(true);
    try {
      await signupService(data);
      setIsSubmitting(false);
      router.push("/", { scroll: false });
    } catch (err) {
      setIsSubmitting(false);
      setError(err.message);
    }
  };

  return (
    <div className="relative h-screen w-1/2 px-20 flex flex-col justify-center items-center">
      <Link className="absolute text-xl top-0 right-0 m-10" href="/login">
        Login
      </Link>
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
        <Button type="submit" className="mt-4" disabled={isSubmitting}>
          {isSubmitting && <Loader2 className="mr-2 h-4 w-4 animate-spin" />}
          Sign up
        </Button>
        <p className="text-destructive text-sm">{error}</p>
      </form>
    </div>
  );
}
