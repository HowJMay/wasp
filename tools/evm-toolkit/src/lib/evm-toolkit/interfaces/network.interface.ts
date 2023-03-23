export interface INetwork {
    id: number;
    text: string;
    apiEndpoint: string;
    faucetEndpoint: string;
    chainAddress: string;
    chainID: number;
    multicallAddress: string;
    blockExplorer?: string;
    networkUrl?: string;
}
