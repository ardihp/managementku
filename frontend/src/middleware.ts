import { paths } from "@/const/routes";
import { NextResponse, NextRequest } from "next/server";

export function middleware(request: NextRequest) {
  const isAuthenticated = !!request.cookies.get("token");

  // If the user is authenticated, continue as normal
  if (isAuthenticated) {
    if (
      request.url.includes(paths.auth.sign_in) ||
      request.url.includes(paths.auth.sign_up)
    ) {
      return NextResponse.redirect(new URL(paths.dashboard.root, request.url));
    }

    return NextResponse.next();
  }

  // Redirect to login page if not authenticated
  if (request.url.includes(paths.dashboard.root)) {
    return NextResponse.redirect(new URL(paths.auth.sign_in, request.url));
  }

  return NextResponse.next();
}

export const config = {
  matcher: ["/dashboard/:path*", "/auth/:path*"],
};
