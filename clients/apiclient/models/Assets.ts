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

import { NativeToken } from '../models/NativeToken';
import { HttpFile } from '../http/http';

export class Assets {
    /**
    * The base tokens (uint64 as string)
    */
    'baseTokens': string;
    'nativeTokens': Array<NativeToken>;
    'nfts': Array<string>;

    static readonly discriminator: string | undefined = undefined;

    static readonly attributeTypeMap: Array<{name: string, baseName: string, type: string, format: string}> = [
        {
            "name": "baseTokens",
            "baseName": "baseTokens",
            "type": "string",
            "format": "string"
        },
        {
            "name": "nativeTokens",
            "baseName": "nativeTokens",
            "type": "Array<NativeToken>",
            "format": ""
        },
        {
            "name": "nfts",
            "baseName": "nfts",
            "type": "Array<string>",
            "format": "string"
        }    ];

    static getAttributeTypeMap() {
        return Assets.attributeTypeMap;
    }

    public constructor() {
    }
}
