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
import {
    EntPaymentEdges,
    EntPaymentEdgesFromJSON,
    EntPaymentEdgesFromJSONTyped,
    EntPaymentEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntPayment
 */
export interface EntPayment {
    /**
     * AccountName holds the value of the "account_name" field.
     * @type {string}
     * @memberof EntPayment
     */
    accountName?: string;
    /**
     * AccountNumber holds the value of the "account_number" field.
     * @type {string}
     * @memberof EntPayment
     */
    accountNumber?: string;
    /**
     * 
     * @type {EntPaymentEdges}
     * @memberof EntPayment
     */
    edges?: EntPaymentEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntPayment
     */
    id?: number;
    /**
     * TransferTime holds the value of the "transfer_time" field.
     * @type {string}
     * @memberof EntPayment
     */
    transferTime?: string;
}

export function EntPaymentFromJSON(json: any): EntPayment {
    return EntPaymentFromJSONTyped(json, false);
}

export function EntPaymentFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntPayment {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'accountName': !exists(json, 'account_name') ? undefined : json['account_name'],
        'accountNumber': !exists(json, 'account_number') ? undefined : json['account_number'],
        'edges': !exists(json, 'edges') ? undefined : EntPaymentEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
        'transferTime': !exists(json, 'transfer_time') ? undefined : json['transfer_time'],
    };
}

export function EntPaymentToJSON(value?: EntPayment | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'account_name': value.accountName,
        'account_number': value.accountNumber,
        'edges': EntPaymentEdgesToJSON(value.edges),
        'id': value.id,
        'transfer_time': value.transferTime,
    };
}

