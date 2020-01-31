package com.standardizer.controller;

import com.standardizer.dtos.Request;
import com.standardizer.dtos.Response;
import com.standardizer.services.Standardizer;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;


@RestController
@RequestMapping("/api/v1/standardize")
public class StandardizerController {

    private final Standardizer standardizer;

    @Autowired
    public StandardizerController(Standardizer standardizer) {
        this.standardizer = standardizer;
    }

    @PostMapping("/")
    public Response standardize(@RequestBody Request request) {
        return standardizer.standardize(request);
    }

}
