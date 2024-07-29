package com.phuc.productservice.response;

import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.Builder;
import lombok.Getter;
import lombok.Setter;

import java.util.ArrayList;
import java.util.List;

@Setter
@Getter
@Builder
public class ResponseError {

    @JsonProperty("message")
    private List<String> message;

    @JsonProperty("error")
    private String error;

    @JsonProperty("statusCode")
    private int statusCode;

    public void addMessage(String error) {
        if (message == null) {
            this.message = new ArrayList<>();
        }

        message.add(error);
    }
}
