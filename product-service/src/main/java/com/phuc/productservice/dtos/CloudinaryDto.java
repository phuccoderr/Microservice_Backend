package com.phuc.productservice.dtos;

import lombok.*;

@Setter
@Getter
@NoArgsConstructor
public class CloudinaryDto {
    private String publicId;
    private String url;

    public CloudinaryDto(String publicId,String url) {
        this.publicId = publicId;
        this.url = url;
    }
}
