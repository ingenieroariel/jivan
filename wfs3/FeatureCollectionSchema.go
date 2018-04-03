package wfs3

// Collected from: http://geojson.org/schema/FeatureCollection.json 2018-04-02
var FeatureCollectionJSONSchema string = `
{
  "$schema": "http://json-schema.org/draft-07/schema#",
  "$id": "http://geojson.org/schema/FeatureCollection.json",
  "title": "GeoJSON FeatureCollection",
  "type": "object",
  "required": [
    "type",
    "features"
  ],
  "properties": {
    "type": {
      "type": "string",
      "emum": [
        "FeatureCollection"
      ]
    },
    "features": {
      "type": "array",
      "items": {
        "title": "GeoJSON Feature",
        "type": "object",
        "required": [
          "type",
          "properties",
          "geometry"
        ],
        "properties": {
          "type": {
            "type": "string",
            "emum": [
              "Feature"
            ]
          },
          "properties": {
            "type": "object"
          },
          "geometry": {
            "oneOf": [
              {
                "type": "null"
              },
              {
                "title": "GeoJSON Point",
                "type": "object",
                "required": [
                  "type",
                  "coordinates"
                ],
                "properties": {
                  "type": {
                    "type": "string",
                    "enum": [
                      "Point"
                    ]
                  },
                  "coordinates": {
                    "type": "array",
                    "minItems": 2,
                    "items": {
                      "type": "number"
                    }
                  },
                  "bbox": {
                    "type": "array",
                    "minItems": 4,
                    "items": {
                      "type": "number"
                    }
                  }
                }
              },
              {
                "title": "GeoJSON LineString",
                "type": "object",
                "required": [
                  "type",
                  "coordinates"
                ],
                "properties": {
                  "type": {
                    "type": "string",
                    "enum": [
                      "LineString"
                    ]
                  },
                  "coordinates": {
                    "type": "array",
                    "minItems": 2,
                    "items": {
                      "type": "array",
                      "minItems": 2,
                      "items": {
                        "type": "number"
                      }
                    }
                  },
                  "bbox": {
                    "type": "array",
                    "minItems": 4,
                    "items": {
                      "type": "number"
                    }
                  }
                }
              },
              {
                "title": "GeoJSON Polygon",
                "type": "object",
                "required": [
                  "type",
                  "coordinates"
                ],
                "properties": {
                  "type": {
                    "type": "string",
                    "enum": [
                      "Polygon"
                    ]
                  },
                  "coordinates": {
                    "type": "array",
                    "items": {
                      "type": "array",
                      "minItems": 4,
                      "items": {
                        "type": "array",
                        "minItems": 2,
                        "items": {
                          "type": "number"
                        }
                      }
                    }
                  },
                  "bbox": {
                    "type": "array",
                    "minItems": 4,
                    "items": {
                      "type": "number"
                    }
                  }
                }
              },
              {
                "title": "GeoJSON MultiPoint",
                "type": "object",
                "required": [
                  "type",
                  "coordinates"
                ],
                "properties": {
                  "type": {
                    "type": "string",
                    "enum": [
                      "MultiPoint"
                    ]
                  },
                  "coordinates": {
                    "type": "array",
                    "items": {
                      "type": "array",
                      "minItems": 2,
                      "items": {
                        "type": "number"
                      }
                    }
                  },
                  "bbox": {
                    "type": "array",
                    "minItems": 4,
                    "items": {
                      "type": "number"
                    }
                  }
                }
              },
              {
                "title": "GeoJSON MultiLineString",
                "type": "object",
                "required": [
                  "type",
                  "coordinates"
                ],
                "properties": {
                  "type": {
                    "type": "string",
                    "enum": [
                      "MultiLineString"
                    ]
                  },
                  "coordinates": {
                    "type": "array",
                    "items": {
                      "type": "array",
                      "minItems": 2,
                      "items": {
                        "type": "array",
                        "minItems": 2,
                        "items": {
                          "type": "number"
                        }
                      }
                    }
                  },
                  "bbox": {
                    "type": "array",
                    "minItems": 4,
                    "items": {
                      "type": "number"
                    }
                  }
                }
              },
              {
                "title": "GeoJSON MultiPolygon",
                "type": "object",
                "required": [
                  "type",
                  "coordinates"
                ],
                "properties": {
                  "type": {
                    "type": "string",
                    "enum": [
                      "MultiPolygon"
                    ]
                  },
                  "coordinates": {
                    "type": "array",
                    "items": {
                      "type": "array",
                      "items": {
                        "type": "array",
                        "minItems": 4,
                        "items": {
                          "type": "array",
                          "minItems": 2,
                          "items": {
                            "type": "number"
                          }
                        }
                      }
                    }
                  },
                  "bbox": {
                    "type": "array",
                    "minItems": 4,
                    "items": {
                      "type": "number"
                    }
                  }
                }
              },
              {
                "title": "GeoJSON GeometryCollection",
                "type": "object",
                "required": [
                  "type",
                  "geometries"
                ],
                "properties": {
                  "type": {
                    "type": "string",
                    "emum": [
                      "GeometryCollection"
                    ]
                  },
                  "geometries": {
                    "type": "array",
                    "items": {
                      "oneOf": [
                        {
                          "title": "GeoJSON Point",
                          "type": "object",
                          "required": [
                            "type",
                            "coordinates"
                          ],
                          "properties": {
                            "type": {
                              "type": "string",
                              "enum": [
                                "Point"
                              ]
                            },
                            "coordinates": {
                              "type": "array",
                              "minItems": 2,
                              "items": {
                                "type": "number"
                              }
                            },
                            "bbox": {
                              "type": "array",
                              "minItems": 4,
                              "items": {
                                "type": "number"
                              }
                            }
                          }
                        },
                        {
                          "title": "GeoJSON LineString",
                          "type": "object",
                          "required": [
                            "type",
                            "coordinates"
                          ],
                          "properties": {
                            "type": {
                              "type": "string",
                              "enum": [
                                "LineString"
                              ]
                            },
                            "coordinates": {
                              "type": "array",
                              "minItems": 2,
                              "items": {
                                "type": "array",
                                "minItems": 2,
                                "items": {
                                  "type": "number"
                                }
                              }
                            },
                            "bbox": {
                              "type": "array",
                              "minItems": 4,
                              "items": {
                                "type": "number"
                              }
                            }
                          }
                        },
                        {
                          "title": "GeoJSON Polygon",
                          "type": "object",
                          "required": [
                            "type",
                            "coordinates"
                          ],
                          "properties": {
                            "type": {
                              "type": "string",
                              "enum": [
                                "Polygon"
                              ]
                            },
                            "coordinates": {
                              "type": "array",
                              "items": {
                                "type": "array",
                                "minItems": 4,
                                "items": {
                                  "type": "array",
                                  "minItems": 2,
                                  "items": {
                                    "type": "number"
                                  }
                                }
                              }
                            },
                            "bbox": {
                              "type": "array",
                              "minItems": 4,
                              "items": {
                                "type": "number"
                              }
                            }
                          }
                        },
                        {
                          "title": "GeoJSON MultiPoint",
                          "type": "object",
                          "required": [
                            "type",
                            "coordinates"
                          ],
                          "properties": {
                            "type": {
                              "type": "string",
                              "enum": [
                                "MultiPoint"
                              ]
                            },
                            "coordinates": {
                              "type": "array",
                              "items": {
                                "type": "array",
                                "minItems": 2,
                                "items": {
                                  "type": "number"
                                }
                              }
                            },
                            "bbox": {
                              "type": "array",
                              "minItems": 4,
                              "items": {
                                "type": "number"
                              }
                            }
                          }
                        },
                        {
                          "title": "GeoJSON MultiLineString",
                          "type": "object",
                          "required": [
                            "type",
                            "coordinates"
                          ],
                          "properties": {
                            "type": {
                              "type": "string",
                              "enum": [
                                "MultiLineString"
                              ]
                            },
                            "coordinates": {
                              "type": "array",
                              "items": {
                                "type": "array",
                                "minItems": 2,
                                "items": {
                                  "type": "array",
                                  "minItems": 2,
                                  "items": {
                                    "type": "number"
                                  }
                                }
                              }
                            },
                            "bbox": {
                              "type": "array",
                              "minItems": 4,
                              "items": {
                                "type": "number"
                              }
                            }
                          }
                        },
                        {
                          "title": "GeoJSON MultiPolygon",
                          "type": "object",
                          "required": [
                            "type",
                            "coordinates"
                          ],
                          "properties": {
                            "type": {
                              "type": "string",
                              "enum": [
                                "MultiPolygon"
                              ]
                            },
                            "coordinates": {
                              "type": "array",
                              "items": {
                                "type": "array",
                                "items": {
                                  "type": "array",
                                  "minItems": 4,
                                  "items": {
                                    "type": "array",
                                    "minItems": 2,
                                    "items": {
                                      "type": "number"
                                    }
                                  }
                                }
                              }
                            },
                            "bbox": {
                              "type": "array",
                              "minItems": 4,
                              "items": {
                                "type": "number"
                              }
                            }
                          }
                        }
                      ]
                    }
                  },
                  "bbox": {
                    "type": "array",
                    "minItems": 4,
                    "items": {
                      "type": "number"
                    }
                  }
                }
              }
            ]
          },
          "bbox": {
            "type": "array",
            "minItems": 4,
            "items": {
              "type": "number"
            }
          }
        }
      }
    },
    "bbox": {
      "type": "array",
      "minItems": 4,
      "items": {
        "type": "number"
      }
    }
  }
}
`
