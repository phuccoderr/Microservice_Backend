package com.phuc.productservice.dtos;

import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

@Setter
@Getter
@NoArgsConstructor
@AllArgsConstructor
public class ProductImageDto {
    @JsonProperty("id")
    private String id;

    @JsonProperty("url")
    private String url;
}
