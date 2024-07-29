package com.phuc.productservice.response;

import com.phuc.productservice.dtos.CategoryDto;
import lombok.Builder;
import lombok.Getter;
import lombok.Setter;

@Setter
@Getter
@Builder
public class ResponseObject {

    private String message;
    private Integer status;
    private Object data;
}
