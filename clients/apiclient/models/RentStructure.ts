/**
 * Wasp API
 * REST API for the Wasp node
 *
 * OpenAPI spec version: 0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { HttpFile } from '../http/http';

export class RentStructure {
    /**
    * The virtual byte cost
    */
    'vByteCost': number;
    /**
    * The virtual byte factor for data fields
    */
    'vByteFactorData': number;
    /**
    * The virtual byte factor for key/lookup generating fields
    */
    'vByteFactorKey': number;

    static readonly discriminator: string | undefined = undefined;

    static readonly attributeTypeMap: Array<{name: string, baseName: string, type: string, format: string}> = [
        {
            "name": "vByteCost",
            "baseName": "vByteCost",
            "type": "number",
            "format": "int32"
        },
        {
            "name": "vByteFactorData",
            "baseName": "vByteFactorData",
            "type": "number",
            "format": "int32"
        },
        {
            "name": "vByteFactorKey",
            "baseName": "vByteFactorKey",
            "type": "number",
            "format": "int32"
        }    ];

    static getAttributeTypeMap() {
        return RentStructure.attributeTypeMap;
    }

    public constructor() {
    }
}
