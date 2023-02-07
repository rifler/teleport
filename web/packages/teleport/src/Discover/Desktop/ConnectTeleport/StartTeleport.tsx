/**
 * Copyright 2023 Gravitational, Inc
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import React from 'react';
import styled from 'styled-components';

import logoSrc from 'design/assets/images/teleport-medallion.svg';

import { Text } from 'design';

import { ButtonPrimary } from 'design/Button';

import {
  StepContent,
  StepInstructions,
  StepTitle,
  StepTitleIcon,
} from 'teleport/Discover/Desktop/ConnectTeleport/Step';

import { usePingTeleport } from 'teleport/Discover/Shared/PingTeleportContext';

interface StartTeleportProps {
  onNext: () => void;
}

interface StepWrapperProps {
  children?: React.ReactNode;
}

function StepWrapper(props: StepWrapperProps) {
  return (
    <StepContent>
      <StepTitle>
        <StepTitleIcon>
          <TeleportIcon />
        </StepTitleIcon>
        4. Start Teleport
      </StepTitle>

      {props.children}
    </StepContent>
  );
}

export function StartTeleport(
  props: React.PropsWithChildren<StartTeleportProps>
) {
  const { active, result, start, timedOut } = usePingTeleport();

  if (timedOut) {
    return (
      <StepWrapper>
        <StepInstructions>
          <Text mb={4}>
            We looked everywhere but we couldn't find your Teleport node.
          </Text>

          <ButtonPrimary disabled={active} onClick={() => start()}>
            Retry
          </ButtonPrimary>
        </StepInstructions>
      </StepWrapper>
    );
  }

  if (result) {
    return (
      <StepWrapper>
        <StepInstructions>
          <Text mb={4}>
            Success! We've detected the new Teleport node you configured.
          </Text>

          <ButtonPrimary onClick={() => props.onNext()}>Next</ButtonPrimary>
        </StepInstructions>
      </StepWrapper>
    );
  }

  return (
    <StepWrapper>
      <StepInstructions>
        <Text mb={4}>Once you've started Teleport, we'll detect it here.</Text>

        <ButtonPrimary disabled={!result} onClick={() => props.onNext()}>
          Next
        </ButtonPrimary>
      </StepInstructions>
    </StepWrapper>
  );
}

const TeleportIcon = styled.div`
  width: 30px;
  height: 30px;
  background: url(${logoSrc}) no-repeat;
  background-size: contain;
  top: 1px;
  position: relative;
`;
