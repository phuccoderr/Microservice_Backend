import { v2 as cloudinary } from 'cloudinary';
import * as process from 'node:process';

export const CloudinaryProvider = {
  provide: 'Cloudinary',
  useFactory: () => {
    return cloudinary.config({
      cloud_name: process.env.CLOUD_NAME,
      api_key: process.env.CLOUD_KEY,
      api_secret: process.env.CLOUD_SECRET
    });
  }
}