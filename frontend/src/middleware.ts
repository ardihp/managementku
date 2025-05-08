import { paths } from "@/const/routes";
import { NextResponse, NextRequest } from "next/server";

export function middleware(request: NextRequest) {
  const isAuthenticated = false;

  // If the user is authenticated, continue as normal
  if (isAuthenticated) {
    return NextResponse.next();
  }

  // Redirect to login page if not authenticated
  return NextResponse.redirect(new URL(paths.auth.sign_in, request.url));
}

export const config = {
  matcher: `/dashboard/:path*`,
};
