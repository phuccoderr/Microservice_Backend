package com.phuc.productservice.service;

import com.phuc.productservice.constants.Constants;
import com.phuc.productservice.dtos.CategoryDto;
import com.phuc.productservice.dtos.CloudinaryDto;
import com.phuc.productservice.exceptions.DataErrorException;
import com.phuc.productservice.exceptions.FuncErrorException;
import com.phuc.productservice.exceptions.ParamValidateException;
import com.phuc.productservice.models.Product;
import com.phuc.productservice.models.ProductImage;
import com.phuc.productservice.repository.ProductImageRepository;
import com.phuc.productservice.repository.ProductRepository;
import com.phuc.productservice.request.RequestProduct;
import com.phuc.productservice.util.FileUploadUtil;
import com.phuc.productservice.util.Utility;
import lombok.RequiredArgsConstructor;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.PageRequest;
import org.springframework.data.domain.Pageable;
import org.springframework.data.domain.Sort;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;
import org.springframework.web.multipart.MultipartFile;

import java.util.List;

@Service
@RequiredArgsConstructor
public class ProductService implements IProductService {
    
    private final ProductRepository proRepository;
    private final ProductImageRepository proImageRepository;
    private final CloudinaryService cloudinaryService;

    @Override
    public Page<Product> getAllProducts(Integer page, Integer limit, String sort, String keyword)
            throws ParamValidateException {

        Utility.checkSortIsAscOrDesc(sort);

        Sort sortDir = Sort.by("name");
        sortDir = sort.equals("asc") ? sortDir.ascending() : sortDir.descending();

        Pageable pageable = PageRequest.of(page - 1,limit,sortDir);

        if (!keyword.isEmpty()) {
            return proRepository.search(keyword, pageable);
        } else  {
            return proRepository.findAll(pageable);
        }
    }

    @Override
    public Product getProduct(String proId) throws DataErrorException {
        return proRepository.findById(proId).orElseThrow( () -> new DataErrorException(Constants.DB_NOT_FOUND) );
    }



    @Override
    public Product createProduct(
            MultipartFile mainImage,
            List<MultipartFile> extraImages,
            RequestProduct requestProduct,
            CategoryDto categoryDto
    ) throws FuncErrorException {

        Product product = new Product();
        setDtoToEntity(product,requestProduct);

        if (mainImage != null && !mainImage.isEmpty()) {
            setMainImage(mainImage,product);
        }
        setExtraImage(extraImages,product);

        if(categoryDto != null) {
            product.setCategoryId(categoryDto.getId());
        }

        if (!requestProduct.getExtraImages().isEmpty()) {
            requestProduct.getExtraImages().forEach(product::addImage
            );
        }
        return proRepository.save(product);
    }

    @Override
    public Product updateProduct(
            Product proInDB,
            RequestProduct requestProduct,
            CategoryDto categoryDto,
            MultipartFile mainImage
    ) throws FuncErrorException {

        setDtoToEntity(proInDB,requestProduct);

        if (mainImage != null && !mainImage.isEmpty()) {
            cloudinaryService.deleteAsyncImage(proInDB.getImageId());
            setMainImage(mainImage,proInDB);
        }

        if(categoryDto != null) {
            proInDB.setCategoryId(categoryDto.getId());
        }

        return proRepository.save(proInDB);
    }

    @Override
    public void deleteProductById(String proId) throws DataErrorException {
        Product product = getProduct(proId);
        proRepository.delete(product);
    }

    @Override
    public void checkNameUnique(String name) throws DataErrorException {
        Product product = proRepository.findByName(name);
        if (product != null) {
            throw new DataErrorException(Constants.DB_ALREADY_EXISTS);
        }
    }

    @Override
    public void checkNameUnique(String oldName, String newName) throws DataErrorException {
        if (!oldName.equals(newName)) {
            checkNameUnique(newName);
        }
    }

    @Override
    public Product addFiles(List<MultipartFile> extraFiles, Product product) throws FuncErrorException {
        for (MultipartFile file : extraFiles) {
            FileUploadUtil.assertAllowed(file, FileUploadUtil.IMAGE_PATTERN);

            String fileName = FileUploadUtil.getFileName(file.getOriginalFilename());

            CloudinaryDto cloudinaryDto = cloudinaryService.uploadImage(file, fileName);

            product.addImage(cloudinaryDto);
        }
        return proRepository.save(product);
    }

    @Override
    @Transactional
    public void deleteFiles(List<String> listFiles, Product product) {
        listFiles.forEach(fileId -> {
            try {
                ProductImage productImage = proImageRepository.findById(fileId).orElseThrow(
                        () -> new DataErrorException(Constants.DB_NOT_FOUND)
                );
                cloudinaryService.deleteAsyncImage(productImage.getImageId());
                proImageRepository.deleteByIdByProductId(productImage.getId(),product.getId());
            } catch (DataErrorException e) {
                throw new RuntimeException(e);
            }
        });
    }

    private void setMainImage(MultipartFile mainFile,Product product)
            throws FuncErrorException {

        FileUploadUtil.assertAllowed(mainFile, FileUploadUtil.IMAGE_PATTERN);

        String fileName = FileUploadUtil.getFileName(mainFile.getOriginalFilename());

        CloudinaryDto result = cloudinaryService.uploadImage(mainFile, fileName);

        product.setImageId(result.getPublicId());
        product.setUrl(result.getUrl());

    }

    private void setExtraImage(List<MultipartFile> extraFile,Product product)
            throws FuncErrorException {

        if (extraFile != null && !extraFile.isEmpty() ) {
            for (MultipartFile file : extraFile) {

                FileUploadUtil.assertAllowed(file, FileUploadUtil.IMAGE_PATTERN);

                String fileName = FileUploadUtil.getFileName(file.getOriginalFilename());

                CloudinaryDto result = cloudinaryService.uploadImage(file, fileName);
                product.addImage(result);
            }

        }
    }

    private void setDtoToEntity(Product product, RequestProduct reqProduct) {
        product.setName(reqProduct.getName());
        product.setAlias(Utility.unAccent(reqProduct.getName()));
        product.setDescription(reqProduct.getDescription());
        product.setStatus(reqProduct.getStatus());
        product.setCost(reqProduct.getCost());
        product.setPrice(reqProduct.getPrice());
        product.setSale(reqProduct.getSale());
        product.setStock(reqProduct.getStock());
    }

}
