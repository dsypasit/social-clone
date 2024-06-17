import { NextResponse, type NextRequest } from "next/server";
import apiClient from "./lib/apiClient";
import { headers } from "next/headers";
import { AxiosError } from "axios";

export default async function middleware(request: NextRequest) {
  const token = request.cookies.get("token")?.value;
  console.log("Middleware processing request:", request.nextUrl.pathname);

  // Handle missing token for non-login routes
  if (!token && !request.nextUrl.pathname.startsWith("/login")) {
    return NextResponse.redirect(new URL("/login", request.url));
  }

  // Validate token with API (assuming a POST request)
  try {
    const res = await apiClient.post(
      "/auth/checktoken",
      {},
      {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      },
    );

    if (res.status !== 204) {
      return NextResponse.redirect(new URL("/login", request.url));
    }
  } catch (err) {
    if (err instanceof AxiosError) {
      if (err?.response?.status == 401) {
        return NextResponse.redirect(new URL("/login", request.url));
      }
    }
    console.error("Error validating token:", err);
    // Consider redirecting to an error page or displaying an appropriate message
    return NextResponse.redirect(new URL("/login", request.url));
  }

  // Token is valid, allow request to proceed
  return NextResponse.next();
}

export const config = {
  matcher: ["/"],
};
