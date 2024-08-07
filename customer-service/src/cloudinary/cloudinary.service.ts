import { Injectable } from '@nestjs/common';
import { v2 as cloudinary } from 'cloudinary';
import { CloudinaryResponse } from './cloudinary-response';
import { CLOUDINARY_CONSTANTS } from '../constants/cloudinary-constants';


const streamifier = require('streamifier');

@Injectable()
export class CloudinaryService {

  uploadFile(file: Express.Multer.File): Promise<CloudinaryResponse> {
    return new Promise<CloudinaryResponse>((resolve, reject) => {
      const uploadStream = cloudinary.uploader.upload_stream(
        {folder: CLOUDINARY_CONSTANTS.FOLDER},
        (error, result) => {
          if (error) return reject(error);
          resolve(result);
        },
      );

      streamifier.createReadStream(file.buffer).pipe(uploadStream);
    });
  }

  deleteImage(publicId: string): Promise<CloudinaryResponse> {
    return new Promise((resolve, reject) => {
      cloudinary.uploader.destroy(publicId, (error, result) => {
        if (error) {
          return reject(error);
        }
        resolve(result);
      });
    });
  }
}
