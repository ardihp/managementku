import { createContext, ReactNode, useState } from "react";
import Cookies from "js-cookie";

interface AuthContext {
  isAuth: boolean;
}

const AuthContext = createContext<AuthContext>({
  isAuth: false,
});

const AuthProvider = ({ children }: { children: ReactNode }) => {
  const [isAuth] = useState(!!Cookies.get("token"));

  return (
    <AuthContext.Provider
      value={{
        isAuth,
      }}
    >
      {children}
    </AuthContext.Provider>
  );
};

export default AuthProvider;
