/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API Playlist Vidoe
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface ControllersRecordinsurance
 */
export interface ControllersRecordinsurance {
    /**
     * 
     * @type {number}
     * @memberof ControllersRecordinsurance
     */
    amountpaidID?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersRecordinsurance
     */
    hospitalID?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersRecordinsurance
     */
    memberID?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersRecordinsurance
     */
    officerID?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersRecordinsurance
     */
    productID?: number;
    /**
     * 
     * @type {string}
     * @memberof ControllersRecordinsurance
     */
    recordinsuranceTime?: string;
}

export function ControllersRecordinsuranceFromJSON(json: any): ControllersRecordinsurance {
    return ControllersRecordinsuranceFromJSONTyped(json, false);
}

export function ControllersRecordinsuranceFromJSONTyped(json: any, ignoreDiscriminator: boolean): ControllersRecordinsurance {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'amountpaidID': !exists(json, 'amountpaidID') ? undefined : json['amountpaidID'],
        'hospitalID': !exists(json, 'hospitalID') ? undefined : json['hospitalID'],
        'memberID': !exists(json, 'memberID') ? undefined : json['memberID'],
        'officerID': !exists(json, 'officerID') ? undefined : json['officerID'],
        'productID': !exists(json, 'productID') ? undefined : json['productID'],
        'recordinsuranceTime': !exists(json, 'recordinsuranceTime') ? undefined : json['recordinsuranceTime'],
    };
}

export function ControllersRecordinsuranceToJSON(value?: ControllersRecordinsurance | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'amountpaidID': value.amountpaidID,
        'hospitalID': value.hospitalID,
        'memberID': value.memberID,
        'officerID': value.officerID,
        'productID': value.productID,
        'recordinsuranceTime': value.recordinsuranceTime,
    };
}


