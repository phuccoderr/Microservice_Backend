package com.phuc.productservice.service;

import com.phuc.productservice.exceptions.DataErrorException;
import com.phuc.productservice.exceptions.ParamValidateException;
import com.phuc.productservice.models.Discount;
import com.phuc.productservice.request.RequestDiscount;
import org.springframework.data.domain.Page;

public interface IDiscountService {

    Discount create(RequestDiscount reqDiscount) throws DataErrorException;
    void deleteById(String id);
    Page<Discount> getAllProductsByCategory(
            Integer page,
            Integer limit,
            String sort,
            String keyword
    ) throws ParamValidateException;
    Discount applyCode(String code) throws DataErrorException;
}
