import { NextResponse, type NextRequest } from "next/server";

export default function middleware(request: NextRequest) {
  const currentUser = request.cookies.get("token")?.value;
  console.log("middleware here");

  if (!currentUser && !request.nextUrl.pathname.startsWith("/login")) {
    return NextResponse.redirect(new URL("/login", request.url));
  }
  return NextResponse.next();
}

export const config = {
  matcher: ["/"],
};
