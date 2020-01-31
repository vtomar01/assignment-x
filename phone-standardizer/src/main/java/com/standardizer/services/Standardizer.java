package com.standardizer.services;

import org.apache.commons.lang.StringUtils;
import org.springframework.stereotype.Service;
import com.standardizer.dtos.Request;
import com.standardizer.dtos.Response;

@Service
public class Standardizer {

    public Response standardize(Request request) {
        Response response = new Response();
        response.setPhone(StringUtils.right(request.getPhone(), 10));
        return response;
    }
}
