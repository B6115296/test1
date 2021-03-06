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
    EntHospitalEdges,
    EntHospitalEdgesFromJSON,
    EntHospitalEdgesFromJSONTyped,
    EntHospitalEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntHospital
 */
export interface EntHospital {
    /**
     * 
     * @type {EntHospitalEdges}
     * @memberof EntHospital
     */
    edges?: EntHospitalEdges;
    /**
     * HospitalName holds the value of the "hospital_name" field.
     * @type {string}
     * @memberof EntHospital
     */
    hospitalName?: string;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntHospital
     */
    id?: number;
}

export function EntHospitalFromJSON(json: any): EntHospital {
    return EntHospitalFromJSONTyped(json, false);
}

export function EntHospitalFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntHospital {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'edges': !exists(json, 'edges') ? undefined : EntHospitalEdgesFromJSON(json['edges']),
        'hospitalName': !exists(json, 'hospital_name') ? undefined : json['hospital_name'],
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntHospitalToJSON(value?: EntHospital | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'edges': EntHospitalEdgesToJSON(value.edges),
        'hospital_name': value.hospitalName,
        'id': value.id,
    };
}


