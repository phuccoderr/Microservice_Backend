package com.phuc.productservice.service;

import com.phuc.productservice.constants.Constants;
import com.phuc.productservice.exceptions.DataErrorException;
import com.phuc.productservice.exceptions.DataNotFoundException;
import com.phuc.productservice.exceptions.ParamValidateException;
import com.phuc.productservice.models.Discount;
import com.phuc.productservice.repository.DiscountRepository;
import com.phuc.productservice.request.RequestDiscount;
import com.phuc.productservice.specifications.DiscountSpecifications;
import com.phuc.productservice.util.Utility;
import lombok.RequiredArgsConstructor;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.PageRequest;
import org.springframework.data.domain.Pageable;
import org.springframework.data.domain.Sort;
import org.springframework.data.jpa.domain.Specification;
import org.springframework.stereotype.Service;

import java.time.LocalDateTime;

@Service
@RequiredArgsConstructor
public class DiscountService implements IDiscountService {
    private final DiscountRepository repository;

    @Override
    public Discount create(RequestDiscount reqDiscount) throws DataErrorException {
        Discount discount = repository.findByCode(reqDiscount.getCode());
        if (discount != null) {
            throw new DataErrorException(Constants.DB_ALREADY_EXISTS);
        }
        Discount newDiscount = Discount.builder()
                .name(reqDiscount.getName())
                .code(reqDiscount.getCode())
                .sale(reqDiscount.getSale())
                .expiryDate(reqDiscount.getExpiryDate())
                .quantity(reqDiscount.getQuantity()).build();

        return repository.save(newDiscount);
    }

    @Override
    public void deleteById(String id) {
        repository.deleteById(id);
    }

    @Override
    public Page<Discount> getAllProductsByCategory(
            Integer page,
            Integer limit,
            String sort,
            String keyword
    ) throws ParamValidateException {
        Utility.checkSortIsAscOrDesc(sort);

        Specification<Discount> spec = Specification.where(null);

        if (!keyword.isEmpty()) {
            spec = spec.and(DiscountSpecifications.withKeyword(keyword));
        }

        Pageable pageable = PageRequest.of(page - 1,limit, Sort.by(
                sort.equals("asc") ? Sort.Direction.ASC : Sort.Direction.DESC,
                "createdAt"
        ));

        return repository.findAll(spec,pageable);

    }

    @Override
    public Discount applyCode(String code) throws DataErrorException {
        Discount discount = repository.findByCode(code);
        if (discount == null) {
            throw new DataErrorException(Constants.DB_NOT_FOUND);
        }

        if (!discount.getExpiryDate().isAfter(LocalDateTime.now())) {
            throw new DataErrorException(Constants.APPLY_DISCOUNT_SUCCESS);
        }

        if (discount.getQuantity() <= 0) {
            throw new DataErrorException(Constants.APPLY_DISCOUNT_SUCCESS);
        }

        discount.setQuantity(discount.getQuantity() - 1);
        return repository.save(discount);
    }

}
