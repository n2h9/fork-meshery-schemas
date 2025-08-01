/* eslint-disable */
/**
 * This file was automatically generated by json-schema-to-typescript.
 * DO NOT MODIFY IT BY HAND. Instead, modify the source JSONSchema file,
 * and run json-schema-to-typescript to regenerate this file.
 */

/**
 * Meshery Connections are managed and unmanaged resources that either through discovery or manual entry are tracked by Meshery. Learn more at https://docs.meshery.io/concepts/logical/connections
 */
export interface HttpsSchemasMesheryIoComponentJson {
  /**
   * ID
   */
  id: string;
  /**
   * Connection Name
   */
  name: string;
  /**
   * Credential ID
   */
  credential_id?: string;
  /**
   * Connection Type
   */
  type: string;
  /**
   * Connection Subtype
   */
  sub_type: string;
  /**
   * Connection Kind
   */
  kind: string;
  metadata?: {
    [k: string]: unknown;
  };
  /**
   * Connection Status
   */
  status:
    | "discovered"
    | "registered"
    | "connected"
    | "ignored"
    | "maintenance"
    | "disconnected"
    | "deleted"
    | "not found";
  /**
   * A Universally Unique Identifier used to uniquely identify entites in Meshery. The UUID core defintion is used across different schemas.
   */
  user_id?: string;
  created_at?: string;
  updated_at?: string;
  deleted_at?: string;
  environments?: HttpsSchemasMesheryIoEnvironmentJson[];
}
/**
 * Meshery Environments allow you to logically group related Connections and their associated Credentials.. Learn more at https://docs.meshery.io/concepts/logical/environments
 */
export interface HttpsSchemasMesheryIoEnvironmentJson {
  /**
   * ID
   */
  id: string;
  /**
   * Environment name
   */
  name: string;
  /**
   * Environment description
   */
  description: string;
  /**
   * Environment organization ID
   */
  organization_id: string;
  /**
   * Environment owner
   */
  owner?: string;
  created_at?: string;
  metadata?: {
    [k: string]: unknown;
  };
  updated_at?: string;
  deleted_at?: string;
}
