package com.phuc.productservice.response;

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
