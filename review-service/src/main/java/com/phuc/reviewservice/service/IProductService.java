package com.phuc.reviewservice.service;

import com.phuc.reviewservice.exeptions.DataErrorException;
import jakarta.servlet.http.HttpServletRequest;

public interface IProductService {
    void findById(String proId, HttpServletRequest request) throws DataErrorException;
}
