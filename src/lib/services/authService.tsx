import { AxiosError } from "axios";
import apiClient from "../apiClient";
import Cookies from "js-cookie";

type LoginError = {
  message: string;
};

interface UserLogin {
  username: string;
  password: string;
}

interface UserSignup {
  username: string;
  password: string;
  email: string;
}

export const loginService = async (credentials: UserLogin) => {
  try {
    const response = await apiClient.post("/auth/login", credentials);
    if (response.data.token) {
      Cookies.set("token", response.data.token, {
        expires: 7, // Expires in 7 days
        path: "/", // Available throughout the entire site
        // domain: "example.com", // Restrict the cookie to a specific domain
        // secure: true, // Only send the cookie over HTTPS
        sameSite: "strict", // Prevent CSRF attacks
      });
    }
    return response;
  } catch (error) {
    if (error instanceof AxiosError) {
      const axiosError = error as AxiosError<LoginError>;
      if (axiosError.response) {
        const { status, data } = axiosError.response;
        if (status === 500) {
          throw new Error("Internal Server Error");
        } else if (status === 404) {
          throw new Error("API Endpoint Not Found");
        } else {
          throw new Error(data.message);
        }
      }
    }
  }
};

export const signupService = async (credentials: UserSignup) => {
  try {
    const response = await apiClient.post("/auth/signup", credentials);
    if (response.data.token) {
      Cookies.set("token", response.data.token, {
        expires: 7, // Expires in 7 days
        path: "/", // Available throughout the entire site
        // domain: "example.com", // Restrict the cookie to a specific domain
        // secure: true, // Only send the cookie over HTTPS
        sameSite: "strict", // Prevent CSRF attacks
      });
    }
    return response;
  } catch (error) {
    if (error instanceof AxiosError) {
      const axiosError = error as AxiosError<LoginError>;
      if (axiosError.response) {
        const { status, data } = axiosError.response;
        if (status === 500) {
          throw new Error("Internal Server Error");
        } else if (status === 404) {
          throw new Error("API Endpoint Not Found");
        } else {
          throw new Error(data.message);
        }
      }
    }
  }
};
