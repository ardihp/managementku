import { paths } from "@/const/routes";
import axios from "axios";
import Cookies from "js-cookie";
import Router from "next/router";
import { toast } from "sonner";

const Api = axios.create({
  baseURL: "http://localhost:8000/api",
});

Api.interceptors.request.use(
  (config) => {
    const token = Cookies.get("token");
    config.headers.Authorization = `Bearer ${token}`;

    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

Api.interceptors.response.use(
  (response) => {
    return response;
  },
  (error) => {
    console.log(error);

    toast(
      error?.response?.data?.message || error?.response?.data || error?.message,
      {
        id: "toast-error",
      }
    );

    if (error.status === 401) {
      Cookies.remove("token");
      Cookies.remove("user");
      Router.push(paths.auth.sign_in);
    }

    return Promise.reject(error);
  }
);

export default Api;
