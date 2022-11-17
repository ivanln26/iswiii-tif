declare namespace NodeJS {
  export interface ProcessEnv {
    readonly BACKEND_URI: string;
    readonly REDIS_URI: string;
  }
}
