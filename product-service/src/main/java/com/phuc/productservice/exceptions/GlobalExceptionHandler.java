package com.phuc.productservice.exceptions;



import com.phuc.productservice.response.ResponseError;
import org.apache.coyote.BadRequestException;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.http.HttpHeaders;
import org.springframework.http.HttpStatus;
import org.springframework.http.HttpStatusCode;
import org.springframework.http.ResponseEntity;
import org.springframework.validation.FieldError;
import org.springframework.web.bind.MethodArgumentNotValidException;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.RestControllerAdvice;
import org.springframework.web.client.HttpClientErrorException;
import org.springframework.web.context.request.WebRequest;
import org.springframework.web.servlet.mvc.method.annotation.ResponseEntityExceptionHandler;

import java.util.List;

@RestControllerAdvice
public class GlobalExceptionHandler extends ResponseEntityExceptionHandler {

    private static final Logger GLOBAL_LOGGER = LoggerFactory.getLogger(GlobalExceptionHandler.class);

    @ExceptionHandler(Exception.class)
    public ResponseEntity<ResponseError> handleGeneralException(Exception ex) {
        return buildErrorResponse(ex, HttpStatus.INTERNAL_SERVER_ERROR);
    }

    @ExceptionHandler(DataErrorException.class)
    public ResponseEntity<ResponseError> handleDataNotFoundException(Exception ex) {

        return buildErrorResponse(ex, HttpStatus.UNPROCESSABLE_ENTITY);
    }

    @ExceptionHandler(FuncErrorException.class)
    public ResponseEntity<ResponseError> handleFuncErrorException(Exception ex) {

        return buildErrorResponse(ex, HttpStatus.BAD_REQUEST);
    }

    @ExceptionHandler(HttpClientErrorException.class)
    public ResponseEntity<ResponseError> handleHttpClientErrorException(Exception ex) {
        return buildErrorResponse(ex, HttpStatus.UNPROCESSABLE_ENTITY);
    }

    @ExceptionHandler(BadRequestException.class)
    public ResponseEntity<ResponseError> handleBadRequestException(Exception ex) {
        return buildErrorResponse(ex, HttpStatus.BAD_REQUEST);
    }

    @ExceptionHandler(ParamValidateException.class)
    public ResponseEntity<ResponseError> handleParamValidateException(Exception ex) {

        return buildErrorResponse(ex, HttpStatus.BAD_REQUEST);
    }

    @Override
    protected ResponseEntity<Object> handleMethodArgumentNotValid(
            MethodArgumentNotValidException ex,
            HttpHeaders headers,
            HttpStatusCode status,
            WebRequest request) {

        GLOBAL_LOGGER.error(ex.getMessage(),ex);

        ResponseError responseError = ResponseError.builder()
                .error(HttpStatus.BAD_REQUEST.getReasonPhrase())
                .statusCode(HttpStatus.BAD_REQUEST.value()).build();

        List<FieldError> fieldError = ex.getBindingResult().getFieldErrors();
        fieldError.forEach( field -> responseError.addMessage(field.getDefaultMessage()));

        return new ResponseEntity<>(responseError, HttpStatus.BAD_REQUEST);
    }

    private ResponseEntity<ResponseError> buildErrorResponse(Exception ex, HttpStatus status) {
        GLOBAL_LOGGER.error(ex.getMessage(),ex);
        ResponseError responseError = ResponseError.builder()
                .statusCode(status.value())
                .error(status.getReasonPhrase())
                .message(List.of(ex.getMessage()))
                .build();
        return new ResponseEntity<>(responseError, status);
    }


}
