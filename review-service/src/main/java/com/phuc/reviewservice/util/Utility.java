package com.phuc.reviewservice.util;

import com.phuc.reviewservice.constants.Constants;
import com.phuc.reviewservice.dtos.ReviewDto;
import com.phuc.reviewservice.exeptions.ParamValidateException;
import com.phuc.reviewservice.models.Review;
import lombok.experimental.UtilityClass;
import org.modelmapper.ModelMapper;

import java.util.List;

@UtilityClass
public class Utility {
    static ModelMapper modelMapper = new ModelMapper();

    public void checkSortIsAscOrDesc(String sort) throws ParamValidateException {
        if (!sort.equals("asc") && !sort.equals("desc")) {
            throw new ParamValidateException(Constants.PARAM_SORT_FAIL);
        }
    }

    public static List<ReviewDto> toListDtos(List<Review> reviews) {
        return reviews.stream().map(Utility::toDto).toList();
    }

    public ReviewDto toDto(Review review) {
        ReviewDto dto = modelMapper.map(review, ReviewDto.class);
        return dto;
    }
}
