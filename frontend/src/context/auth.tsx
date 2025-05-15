import {
  createContext,
  Dispatch,
  ReactNode,
  SetStateAction,
  useEffect,
  useState,
} from "react";
import Cookies from "js-cookie";

interface AuthContext {
  isAuth: boolean;
  setIsAuth: Dispatch<SetStateAction<boolean>>;
}

export const AuthContext = createContext<AuthContext>({
  isAuth: false,
  setIsAuth: () => {},
});

const AuthProvider = ({ children }: { children: ReactNode }) => {
  const [isAuth, setIsAuth] = useState(!!Cookies.get("token"));

  useEffect(() => {
    const handleTokenChange = () => {
      setIsAuth(!!Cookies.get("token"));
    };

    window.addEventListener("storage", handleTokenChange);

    return () => {
      window.addEventListener("storage", handleTokenChange);
    };
  }, []);

  return (
    <AuthContext.Provider value={{ isAuth, setIsAuth }}>
      {children}
    </AuthContext.Provider>
  );
};

export default AuthProvider;
