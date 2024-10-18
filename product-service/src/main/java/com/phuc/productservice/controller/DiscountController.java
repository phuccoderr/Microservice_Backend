package com.phuc.productservice.controller;


import com.phuc.productservice.constants.Constants;
import com.phuc.productservice.dtos.DiscountDto;
import com.phuc.productservice.dtos.PaginationDto;
import com.phuc.productservice.exceptions.DataErrorException;
import com.phuc.productservice.exceptions.ParamValidateException;
import com.phuc.productservice.models.Discount;
import com.phuc.productservice.request.RequestDiscount;
import com.phuc.productservice.response.ResponseObject;
import com.phuc.productservice.service.DiscountService;
import com.phuc.productservice.util.Utility;
import jakarta.validation.Valid;
import lombok.RequiredArgsConstructor;
import org.springframework.data.domain.Page;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping("/api/v1/discounts")
@RequiredArgsConstructor
public class DiscountController {

    private final DiscountService service;

    @GetMapping
    public ResponseEntity<ResponseObject> listByPage(
            @RequestParam(value = "page", defaultValue = "1") Integer page,
            @RequestParam(value = "limit", defaultValue = "10") Integer limit,
            @RequestParam(value = "sort", defaultValue = "asc") String sort,
            @RequestParam(value = "keyword", defaultValue = "") String keyword) throws ParamValidateException {
        Page<Discount> pages = service.getAllProductsByCategory(page, limit, sort, keyword);
        List<Discount> discounts = pages.getContent();
        List<DiscountDto> listDtos = Utility.toListDiscountDtos(discounts);

        PaginationDto paginationDto = new PaginationDto(pages,listDtos);

        return new ResponseEntity<>(ResponseObject.builder()
                .status(HttpStatus.OK.value())
                .message(Constants.GET_ALL_SUCCESS)
                .data(paginationDto).build(), HttpStatus.OK);
    }

    @PostMapping
    public ResponseEntity<ResponseObject> create(@RequestBody() @Valid RequestDiscount discount) throws DataErrorException {
        Discount newDiscount = service.create(discount);

        DiscountDto dto = Utility.toDiscountDto(newDiscount);

        return new ResponseEntity<>(ResponseObject.builder()
                .status(HttpStatus.OK.value())
                .message(Constants.CREATE_DISCOUNT_SUCCESS)
                .data(dto).build(), HttpStatus.OK);
    }

    @DeleteMapping("/{id}")
    public ResponseEntity<ResponseObject> delete(@PathVariable("id") String id) {
         service.deleteById(id);

        return new ResponseEntity<>(ResponseObject.builder()
                .status(HttpStatus.OK.value())
                .message(Constants.DELETE_DISCOUNT_SUCCESS)
                .data("").build(), HttpStatus.OK);
    }

    @PostMapping("/apply/{code}")
    public ResponseEntity<ResponseObject> applyCode(@PathVariable("code") String code) throws DataErrorException {
        Discount discount = service.applyCode(code);
        DiscountDto discountDto = Utility.toDiscountDto(discount);

        return new ResponseEntity<>(ResponseObject.builder()
                .status(HttpStatus.OK.value())
                .message(Constants.DELETE_DISCOUNT_SUCCESS)
                .data(discountDto).build(), HttpStatus.OK);
    }

    @GetMapping("/{code}")
    public ResponseEntity<ResponseObject> get(@PathVariable("code") String code) throws DataErrorException {
        Discount discount = service.getCode(code);
        DiscountDto discountDto = Utility.toDiscountDto(discount);

        return new ResponseEntity<>(ResponseObject.builder()
                .status(HttpStatus.OK.value())
                .message(Constants.DELETE_DISCOUNT_SUCCESS)
                .data(discountDto).build(), HttpStatus.OK);
    }

}
