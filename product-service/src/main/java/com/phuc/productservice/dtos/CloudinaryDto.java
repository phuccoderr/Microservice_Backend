package com.phuc.productservice.dtos;

import lombok.Builder;
import lombok.Getter;
import lombok.Setter;

@Setter
@Getter
@Builder
public class CloudinaryDto {
    private String publicId;
    private String url;
}
